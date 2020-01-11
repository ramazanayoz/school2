package collections;

public class Student {

	private String name;
	private String surname;
	private int id;
	
	
	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getSurname() {
		return surname;
	}

	public void setSurname(String surname) {
		this.surname = surname;
	}

	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public Student(String name, String surname, int id) {
		super();
		this.name = name;
		this.surname = surname;
		this.id = id;
	}
	
	public String toString(){
		return "id=" + id +", name=" + name + ", surname=" + surname;
	}
	
	public boolean equals(Object obj){
		if (obj instanceof Student){
			return id ==((Student)obj).id;
		}
		return false;
	}
	
	public int hashCode(){
		return id;
	}


	
}
