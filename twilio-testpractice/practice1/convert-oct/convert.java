private static String convertIntToOct(int num) {
		
		char[] c = new char[]{'0','1','2','3','4','5','6','7'};
		
		StringBuilder builder = new StringBuilder();
		while (num > 0) {
			int oct = num % 8;
			builder.append(c[oct]);
			num = num / 8;
		}
		return builder.reverse().toString();
	}
