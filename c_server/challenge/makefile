main: main.c state.o
	gcc main.c state.o -o main

state.o: state.c
	@COUNT=$$(wc -c < state.c); \
	if [ $$COUNT -gt 80 ]; then \
		echo "ERROR: File state.c breaks computer restrictions and is too long."; \
		echo "Must be less than 80 characters. Currently the code is $$COUNT characters."; \
		exit 1; \
	fi; \
	if grep -q 'i' state.c; then \
		echo "ERROR: File state.c contains forbidden character 'i'."; \
		echo "This character is considered 'complex' and is disallowed on this architecture"; \
		exit 1; \
	fi; \
	if grep -q '\[' state.c; then \
		echo "ERROR: File state.c contains forbidden character '['."; \
		echo "This character is considered 'complex' and is disallowed on this architecture"; \
		exit 1; \
	fi; \
	if grep -q '\]' state.c; then \
		echo "ERROR: File state.c contains forbidden character ']'."; \
		echo "This character is considered 'complex' and is disallowed on this architecture"; \
		exit 1; \
	fi; \
		if grep -q '*' state.c; then \
		echo "ERROR: File state.c contains forbidden character '*'."; \
		echo "This character is considered 'complex' and is disallowed on this architecture"; \
		exit 1; \
	fi; \
	gcc state.c -c