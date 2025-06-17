import serial

PORT = "/dev/ttyUSB0"
def update_score(ser):
    try:
        with open("data/score.txt", "a") as f:
            while True:
                line = ser.readline()
                if line:
                    line = line.decode("utf-8").strip()

                    print("New data:", line)
                    f.write(line + "\n")
    except KeyboardInterrupt:
        print("\nTerminated by user")
    except Exception as e:
        print("Error in opening file:")
        print(e)
    finally:
        close_serial(ser)

def close_serial(ser):
    if ser.is_open:
        ser.close()
    print("Serial port is closed")
def open_serial():
    ser = -1
    try:
        ser = serial.Serial(PORT, 9600, timeout=1)
        if ser.is_open:
            print("Serial is open. Receiving data...")  
    except serial.SerialException as e:
        print("Error in opening serial port:")
        print(e)
    return ser

def main():
    ser = open_serial()
    if ser != -1: update_score(ser)

if __name__ == "__main__": main()
