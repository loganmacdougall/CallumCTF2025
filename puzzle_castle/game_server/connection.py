import select
import socket
import struct
import threading

class ConnectionSocket():
  def __init__(self, host: str, port: int):
    self._closed = False
    server_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    server_sock.bind((host, port))
    server_sock.listen()

    self.socket = server_sock

  def __enter__(self):
      return self

  def __exit__(self, *args):
      if not self._closed:
          self.close()

  def close(self):
    self.socket.close()
    self._closed = True

  def accept(self):
    conn, addr = self.socket.accept()
    conn.setsockopt(socket.SOL_SOCKET, socket.SO_KEEPALIVE, 1)

    # Linux-specific settings
    conn.setsockopt(socket.IPPROTO_TCP, socket.TCP_KEEPIDLE, 30)
    conn.setsockopt(socket.IPPROTO_TCP, socket.TCP_KEEPINTVL, 10)
    conn.setsockopt(socket.IPPROTO_TCP, socket.TCP_KEEPCNT, 3)

    return Connection(conn)

class Connection():
  def __init__(self, conn: socket.socket):
    self._closed = False
    self.conn = conn
    self.conn.settimeout(0.2)
    self.send_lock = threading.Lock()

  def __enter__(self):
      return self

  def __exit__(self, *args):
      if not self._closed:
          self.close()
  
  def close(self):
    if not self._closed:
      self.conn.close()
      self._closed = True

  def is_closed(self) -> bool:
    return self._closed

  def bytes_available(self) -> bool:
    s, _, _ = select.select([self.conn], [], [], 0)
    
    if len(s) > 0:
      try:
        data = self.conn.recv(1, socket.MSG_PEEK)
      except Exception as e:
        return False

      if data:
        return True
      else:
        self.close()
    
    return False
  
  def wait_for_data(self) -> bool:
    s, _, _ = select.select([self.conn], [], [])
    
    if len(s) > 0:
      try:
        data = self.conn.recv(1, socket.MSG_PEEK)
      except Exception as e:
        return False

      if data:
        return True
      else:
        self.close()
    
    return False

  def recv(self) -> bytes | None:
    length_bytes = self.conn.recv(4)
    if not length_bytes:
        return None
    length = struct.unpack('<I', length_bytes)[0]
    data = b''
    while len(data) < length:
        packet = self.conn.recv(length - len(data))
        if not packet:
            return None
        data += packet
    return data

  def send(self, data: bytes):
    length = struct.pack('<I', len(data))
    with self.send_lock:
      try:
        self.conn.sendall(length + data)
      except Exception as e:
         pass

  def send_json(self, obj: dict):
    import json
    data = json.dumps(obj).encode('utf-8')
    self.send(data)