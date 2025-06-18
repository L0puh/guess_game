import serial

PORT = "/dev/ttyUSB0"
FILE = "data/score.txt"

def read_result():
    try:
        with open(FILE, "r") as f:
            line = f.readline().strip()
            if line and line.isdigit():
                return int(line)
            return 0
    except Exception as e:
        print("Error in opening file:")
        print(e)

def write_result(res):
    try:
        with open(FILE, "w") as f:
            f.write(str(res))
    except Exception as e:
        print("Error in opening file:")
        print(e)

def update_score(ser):
    try:
        while True:
            line = ser.readline()
            if line:
                line = line.decode("utf-8").strip()
                res = read_result()
                if line.isdigit():
                    res += int(line)
                else: print("NO DIGIT: ", line)
                print("new res:", res)
                write_result(res)

    except KeyboardInterrupt:
        print("\nTerminated by user")
    except Exception as e:
        print("ERROR: ", e)
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
    except Exception as e:
        print("ERROR: ", e)
    return ser

def main():
    ser = open_serial()
    if ser != -1: update_score(ser)

if __name__ == "__main__": main()
