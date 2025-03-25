package main

func main() {
	// keyblock := "D0112P0AN00N0000EECE9D7B08076495AAB9060B4FF22CB9007299ADF83E1CADAD3FAE2F97C103E664B4514A41B7D117C43EFDEF66932810"

	// kb, err := tr31.ParseRawKeyBlock(keyblock, 8)
	// cobra.CheckErr(err)

	// pretty.Println("kb", kb)

	// x := aes.NewCipher(keyComponent1)
	// NewCBCMAC
	// // Mac mac = Mac.getInstance("AESCMAC", new BouncyCastleProvider());
	// var keyComponent1 []byte = hex.MustParseString("0731B84D0170C7B10A53AF222B21C818A87EF85BCE05ABB2EE7079973C874702");
	// var nullBytes []byte = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00};
	// secretKeySpec1 := new SecretKeySpec(keyComponent1, "AES");
	// mac.init(secretKeySpec1);
	// var macData1 []byte = mac.doFinal(nullBytes);
	// System.out.println("key component1 KCV :: " + HexUtil.bytesToHex(macData1).substring(0, 6));
}

// public static void main(String args[]) throws NoSuchAlgorithmException, InvalidKeyException, NoSuchPaddingException, BadPaddingException, IllegalBlockSizeException, InvalidAlgorithmParameterException {
// 	Mac mac = Mac.getInstance("AESCMAC", new BouncyCastleProvider());
// 	var keyComponent1 []byte = hex.MustParseString("0731B84D0170C7B10A53AF222B21C818A87EF85BCE05ABB2EE7079973C874702");
// 	var nullBytes []byte = new byte[]{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00};
// 	SecretKeySpec secretKeySpec1 = new SecretKeySpec(keyComponent1, "AES");
// 	mac.init(secretKeySpec1);
// 	var macData1 []byte = mac.doFinal(nullBytes);
// 	System.out.println("key component1 KCV :: " + HexUtil.bytesToHex(macData1).substring(0, 6));

// 	var keyComponent2 []byte = hex.MustParseString("F384E2721F696AF08CB03FD05B768B8542DED0CDBBFEB392F8F77E6707EA99FD");
// 	SecretKeySpec secretKeySpec2 = new SecretKeySpec(keyComponent2, "AES");
// 	mac.init(secretKeySpec2);
// 	var macData2 []byte = mac.doFinal(nullBytes);
// 	System.out.println("key component2 KCV :: " + HexUtil.bytesToHex(macData2).substring(0, 6));

// 	var zcmkBytes []byte = byteXOR(keyComponent1, keyComponent2);

// 	SecretKeySpec secretKeySpec = new SecretKeySpec(zcmkBytes, "AES");
// 	mac.init(secretKeySpec);
// 	var macData []byte = mac.doFinal(nullBytes);
// 	System.out.println("key KCV :: " + HexUtil.bytesToHex(macData).substring(0, 6));
// 	System.out.println("zcmk key :: " + HexUtil.bytesToHex(zcmkBytes));

// //        byte[] zcmkBytes = hex.MustParseString("F4B55A3F1E19AD4186E390F27057439DEAA0289675FB1820168707F03B6DDEFF");

// 	var encKeyDerData1 []byte = new byte[]{0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x01, 0x00};
// 	var encKeyDerData2 []byte = new byte[]{0x02, 0x00, 0x00, 0x00, 0x00, 0x04, 0x01, 0x00};
// 	var macKeyDerData1 []byte = new byte[]{0x01, 0x00, 0x01, 0x00, 0x00, 0x04, 0x01, 0x00};
// 	var macKeyDerData2 []byte = new byte[]{0x02, 0x00, 0x01, 0x00, 0x00, 0x04, 0x01, 0x00};

// 	SecretKeySpec zcmkSecretKey = new SecretKeySpec(zcmkBytes, "AES");
// 	mac.init(zcmkSecretKey);
// 	var encKeyPart1 []byte = mac.doFinal(encKeyDerData1);
// 	System.out.println("Enc Key part 1 :: " + HexUtil.bytesToHex(encKeyPart1));

// 	var encKeyPart2 []byte = mac.doFinal(encKeyDerData2);
// 	System.out.println("Enc Key part 2 :: " + HexUtil.bytesToHex(encKeyPart2));
// 	var encKey []byte = new byte[encKeyPart1.length + encKeyPart2.length];
// 	System.arraycopy(encKeyPart1, 0, encKey, 0, encKeyPart1.length);
// 	System.arraycopy(encKeyPart2, 0, encKey, encKeyPart1.length, encKeyPart2.length);
// 	System.out.println("Enc Key is :: "+ HexUtil.bytesToHex(encKey));

// 	SecretKeySpec encKeySpec = new SecretKeySpec(encKey, "AES");
// 	var iv []byte = hex.MustParseString("BDE362D70098801017CD766C9BF72CD0");

// 	//RD0144P0AN00S0000
// 	// 774644A3C0FBDC1B8F6B1B288D1730D61F5921B988B3825B5900F3A1B370B6CA49CD4D6728E9BEBF63B9119374D07C66
// 	// BDE362D70098801017CD766C9BF72CD0

// 	var confidentialData []byte = hex.MustParseString("774644A3C0FBDC1B8F6B1B288D1730D61F5921B988B3825B5900F3A1B370B6CA49CD4D6728E9BEBF63B9119374D07C66");
// 	var decryptedBytes []byte = decryptWithAESCBC(confidentialData, encKeySpec, iv);
// 	System.out.println("Decrypted key data :: "+ HexUtil.bytesToHex(decryptedBytes));

// 	var keyBytes []byte = new byte[32];
// 	System.arraycopy(decryptedBytes, 2, keyBytes, 0, 32 );
// 	System.out.println("Clear key is :: "+ HexUtil.bytesToHex(keyBytes));

// 	SecretKeySpec zpkKey = new SecretKeySpec(keyBytes, "AES");
// 	mac.init(zpkKey);
// 	var kcvDataZPK []byte = mac.doFinal(nullBytes);

// 	System.out.println("KCV data for ZPK :: " + HexUtil.bytesToHex(kcvDataZPK));
// 	System.out.println("KCV for ZPK  :: " + HexUtil.bytesToHex(kcvDataZPK).substring(0,6));

// }

// public static byte[] byteXOR(byte[] input1, byte[] input2) {
// 	var output []byte = new byte[input1.length];
// 	for (int i = 0; i < output.length; i++) {
// 		output[i] = (byte) (input1[i] ^ input2[i]);
// 	}
// 	return output;
// }

// public static byte[] decryptWithAESCBC(byte[] cipherData, SecretKey key, byte[] iv) throws NoSuchPaddingException, NoSuchAlgorithmException, InvalidKeyException, BadPaddingException, IllegalBlockSizeException, InvalidAlgorithmParameterException {
// 	var clearData []byte = null;
// 	Cipher cipher = Cipher.getInstance("AES/CBC/NoPadding");
// 	IvParameterSpec encIv = new IvParameterSpec(iv, 0, iv.length);
// 	cipher.init(Cipher.DECRYPT_MODE, key, encIv);
// 	clearData = cipher.doFinal(cipherData);
// 	return clearData;
// }
