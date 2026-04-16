
void setup() {
  // Configurações iniciais
  pinMode(LED_BUILTIN, OUTPUT);
}

void loop() {
  // Código que se repete para sempre
  digitalWrite(LED_BUILTIN, HIGH);
  delay(1000);
  digitalWrite(LED_BUILTIN, LOW);
  delay(1000);
}
