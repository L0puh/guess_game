int LED_YELLOW = 5;
int LED_GREEN = 6;
int LED_RED = 7;

int B_YELLOW = 12;
int B_GREEN = 11;
int B_RED = 10;


void setup() {
  Serial.begin(9600);
  pinMode(LED_RED,    OUTPUT);
  pinMode(LED_YELLOW, OUTPUT);
  pinMode(LED_GREEN,  OUTPUT);

  pinMode(B_RED,    INPUT);
  pinMode(B_YELLOW, INPUT);
  pinMode(B_GREEN,  INPUT);

  randomSeed(analogRead(A0));
}

void generate_random_led(int n, int* arr) {
  int prev = -1;
  for (int i = 0; i < n; i++){
    int r = random(1, 4);
    while (r == prev) 
      r = random(1, 4);
    prev = r;
    arr[i] = r;
  }
}
void clear() {
  digitalWrite(LED_RED, LOW);
  digitalWrite(LED_GREEN, LOW);
  digitalWrite(LED_YELLOW, LOW);
}

void show_lights(int* arr, int n, int SPEED) {
  for (int i = 0; i < n; i++){
    int r = arr[i];
    switch(r) {
      case 1: 
        digitalWrite(LED_YELLOW, HIGH);
        delay(SPEED);
        digitalWrite(LED_YELLOW, LOW);
        break;
      case 2: 
        digitalWrite(LED_GREEN, HIGH);
        delay(SPEED);
        digitalWrite(LED_GREEN, LOW);
        break;
      case 3: 
        digitalWrite(LED_RED, HIGH);
        delay(SPEED);
        digitalWrite(LED_RED, LOW);
        break;
    }
  }
}
void read_input(int *input, int n) {
  int i =0;
  int prev = -1;
  while (i < n){
    int red = digitalRead(B_RED);
    int yellow = digitalRead(B_YELLOW);
    int green = digitalRead(B_GREEN);

    if (yellow && prev != 1) {
        input[i++] = 1;
        prev = 1;
    }
    if (green && prev != 2){
        input[i++] = 2;
        prev = 2;
    }
    if (red && prev != 3){
        input[i++] = 3;
        prev = 3;
      }
    delay(100); 
  }
}

bool check_input(int *input, int *arr, int n){
  for (int i = 0; i < n; i++){
    if (input[i] != arr[i]){
      return false;
    }
  }
  return true;
}
void loop() {
  while (Serial.available() <= 0){
      ;
  }
  int REP = 4;
  int SPEED = 100;
  String message = Serial.readStringUntil('\n');
  int indx = message.indexOf(',');
  if (indx > 0){
    String speed_str = message.substring(0, indx);
    String repeat_str = message.substring(indx+1);
    SPEED = speed_str.toInt();
    REP = repeat_str.toInt();
  }
  Serial.print("RECIEVED: ");
  Serial.println(message);
  int arr[REP], input[REP];
  clear();
  generate_random_led(REP, arr);
  show_lights(arr, REP, SPEED);
  read_input(input, REP);
 
  bool res = check_input(input, arr, REP);
  write_result(res);
  
}

void write_result(int res) {
  if (Serial){
    Serial.println(res);
    delay(1000);
  }
}

