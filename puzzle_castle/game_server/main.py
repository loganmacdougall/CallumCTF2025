# server.py
from connection import ConnectionSocket, Connection
from game_runner import GameRunner
from ast_validator import analyze_code
from worlds.world1 import World1
from worlds.world2 import World2
from worlds.world3 import World3
from worlds.world4 import World4
from worlds.world5 import World5 
from worlds.worldF import WorldF 
import os
import json
import base64
import time
import threading

HOST = "0.0.0.0"
PORT = 8083

worlds = [
  World1(), World2(), World3(), World4(), World5(), WorldF()
]

def run_game(conn: Connection):
    with conn:
      try:
        peername = conn.conn.getpeername()
      except Exception:
        peername = "<unknown>"
      print(f"[THREAD {threading.get_ident()}] New connection established: {peername}")

      conn.wait_for_data()

      if conn.is_closed():
        print(f"[THREAD {threading.get_ident()}] Connection closed by peer: {peername}")
        return

      bytes = conn.recv()
      data = json.loads(bytes.decode('utf-8'))

      if "code" not in data:
        print(f"[THREAD {threading.get_ident()}] No code provided by peer: {peername}")
        return

      try:       
        code = base64.b64decode(data["code"]).decode('utf-8')
        tree = analyze_code(code)
        game = GameRunner(worlds, tree)
      except Exception as e:
        response = {
          "status": "error",
          "console": [str(e)]
        }
        conn.send_json(response)
        print(f"[THREAD {threading.get_ident()}] Code analysis/compilation failed for peer {peername}: {e}")
        return

      start_time = time.time()
      
      update_thread = threading.Thread(target=send_regular_updates, args=(conn,game,start_time), daemon=True)
      update_thread.start()

      while True:
        if not conn.wait_for_data():
          print(f"[THREAD {threading.get_ident()}] Connection closed by peer: {peername}")
          break
        
        bytes = conn.recv()
        data = json.loads(bytes.decode('utf-8'))

        for key, value in data.items():
          if key == "action":
            try:
              game.run_action(value)
            except Exception as e:
              response = {
                "status": "error",
                "console": [str(e)]
              }
              conn.send_json(response)
        time.sleep(0.01)

def send_regular_updates(conn: Connection, game: GameRunner, start_time: float):
    last_time_since_update = 0
    refresh_rate = 1 / 30

    while True:
      if conn.is_closed():
        break

      if time.time() - last_time_since_update < refresh_rate:
        continue

      response = {
        "status": "update",
        "state": base64.b64encode(game.worlds[game.current_world].export_view()).decode('utf-8'),
        "time": max(60 - (time.time() - start_time), 0),
      }

      if len(game.console) > 0:
        response["console"] = [t for t in game.console]
        game.console = []

      if len(game.text) > 0:
        response["text"] = [t for t in game.text]
        game.text = []

      conn.send_json(response)
      last_time_since_update = time.time()

      if game.worlds[game.current_world].game_over:
         break

      time.sleep(0.01)

def main():
    print(f"Setting up server with for {HOST}:{PORT}")
    with ConnectionSocket(HOST, PORT) as server:
        print(f"Server listening on {HOST}:{PORT}")
        while True:
            conn = server.accept()

            # start worker thread that owns the connection
            thread = threading.Thread(target=run_game, args=(conn,), daemon=True)
            thread.start()
            print(f"[MAIN] started thread {thread.ident}; active threads: {threading.active_count()}")

if __name__ == "__main__":
  main()
