"""
This module contains functions for connecting to a WebSocket server and sending/receiving messages.
"""

import json
import websocket

def on_message(ws, message):
    """
    Handle incoming messages from the server.

    Args:
        ws: The WebSocket connection object.
        message: The message that was received.

    Returns:
        None.
    """
    print(f"Received message: {message}")

def on_error(ws, error):
    """
    Handle connection errors.

    Args:
        ws: The WebSocket connection object.
        error: The error that occurred.

    Returns:
        None.
    """
    print(f"Connection error: {error}")

def on_close(ws):
    """
    Handle connection closure.

    Args:
        ws: The WebSocket connection object.

    Returns:
        None.
    """
    print("Connection closed")

def on_open(ws):
    """
    Handle connection opening.

    Args:
        ws: The WebSocket connection object.

    Returns:
        None.
    """
    print("Connection opened")
    # Send a message to the server
    message = {
        "type": "chat",
        "data": "Hello, server!"
    }
    ws.send(json.dumps(message))

if __name__ == "__main__":
    # Establish a WebSocket connection to the server
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp("ws://localhost:8000/",
                                on_open=on_open,
                                on_message=on_message,
                                on_error=on_error,
                                on_close=on_close)
    ws.run_forever()
