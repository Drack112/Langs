fn main(){

    // signed = Positivos e negativos
    // todos os signeds sempre vão começar com 0 no positivo
    // unsigned = Apenas positivos

    let signed_8 :i8 = -128; 
    let unsigned_8 :u8 = 127; 

    let signed_16 :i16 = -32768; 
    let unsigned_16 :u16 = 32767;
   
    let signed_32 :i32 = -2147483648;
    let unsigned_32 :u32 = 2147483647;

    let signed_64 :i64 = -9223372036854775808;
    let unsigned_64 :u64 = 	9223372036854775807;

    let signed_128 :i128 = -170141183460469231731687303715884105728;
    let unsigned_128 :u128 = 170141183460469231731687303715884105727;
    
    println!("Signed 8: {}", signed_8);
    println!("Unsigned 8: {}", unsigned_8);

    println!("Signed 16: {}", signed_16);
    println!("Unsigned 16: {}", unsigned_16);

    println!("Signed 32: {}", signed_32);
    println!("Unsigned 32: {}", unsigned_32);

    println!("Signed 64: {}", signed_64);
    println!("Unsigned 64: {}", unsigned_64);

    println!("Signed 128: {}", signed_128);
    println!("Unsigned 128: {}", unsigned_128);
}