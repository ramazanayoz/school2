package collections;

import java.util.Comparator;

public class StringComparator implements Comparator<String> {

	@Override
	public int compare(String o1, String o2) {
		// TODO Auto-generated method stub
		return o1.toLowerCase().compareTo(o2.toLowerCase());
	}

}
