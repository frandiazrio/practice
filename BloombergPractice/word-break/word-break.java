public boolean wordBreak(String s, List<String> wordDict) {
        
    Set<String> dictionary = new HashSet<>(wordDict);
    
    return canSegmentString(s, dictionary);
}

private boolean canSegmentString(String s, Set<String> dictionary) {
    for(int i = 1; i <= s.length(); i++) {
       String first = s.substring(0, i);
       if(dictionary.contains(first)) {
           String second = s.substring(i, s.length());

           if(second == null ||
              second.length() == 0 ||
              dictionary.contains(second) ||
              canSegmentString(second, dictionary)) {
               return true;
           }
       }
   }
   return false;
}    