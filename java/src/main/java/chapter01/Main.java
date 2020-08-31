package chapter01;

public class Main {
    public static void main(String[] args) {
        Customer customer = new Customer("hara");

        Movie movie1 = new Movie("通常", Movie.REGULAR);
        Rental rental1 = new Rental(movie1, 3);
        customer.addRenatals(rental1);

        Movie movie2 = new Movie("新作", Movie.NEW_RELEASE);
        Rental rental2 = new Rental(movie2, 5);
        customer.addRenatals(rental2);

        System.out.println(customer.statement());
        System.out.println(customer.htmlStatement());
    }
}
