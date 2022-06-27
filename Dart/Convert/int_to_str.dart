main() {
  String oneAsString = 1.toString();
  assert(oneAsString == '1');

  String piAsString = 3.131341441.toStringAsFixed(2);
  print(piAsString);
}
