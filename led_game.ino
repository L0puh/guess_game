int LED_YELLOW = 5;
int LED_GREEN = 6;
int LED_RED = 7;

int B_YELLOW = 12;
int B_GREEN = 11;
int B_RED = 10;

int SPEED = 800;

void setup() {
  Serial.begin(9600);
  pinMode(LED_RED,    OUTPUT);
  pinMode(LED_YELLOW, OUTPUT);
  pinMode(LED_GREEN,  OUTPUT);

  pinMode(B_RED,    INPUT);
  pinMode(B_YELLOW, INPUT);
  pinMode(B_GREEN,  INPUT);

  randomSeed(12345);
}

void generate_random_led(int n, int* arr) {
  int prev = -1;
  for (int i = 0; i < n; i++){
    int r = random(1, 4);
    while (r == prev) 
      r = random(1, 4);
    prev = r;
    arr[i] = r;
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
void clear() {
  digitalWrite(LED_RED, LOW);
  digitalWrite(LED_GREEN, LOW);
  digitalWrite(LED_YELLOW, LOW);
}
void loop() {
  int b_red = digitalRead(B_RED);
  int b_green = digitalRead(B_GREEN);
  int b_yellow = digitalRead(B_YELLOW);
  
  int n = 5;
  int arr[n];
  generate_random_led(n, arr);
  clear(); 

  for (int i = 0; i < n; i++){
    Serial.print(arr[i]);
  }
  Serial.println();
  delay(100000);
}



