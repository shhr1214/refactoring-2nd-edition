package chapter01;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

class CustomerTest {
    // テスト結果が正しいかどうかは知らないけど、少なくとも
    // リファクタリング前と同じ動きをしていれば良いでしょう。
    @Test
    @DisplayName("修正後も Main.java の実行結果と同じ結果を返している")
    void test() {
        Customer customer = new Customer("hara");

        Movie movie1 = new Movie("通常", Movie.REGULAR);
        Rental rental1 = new Rental(movie1, 3);
        customer.addRenatals(rental1);

        Movie movie2 = new Movie("新作", Movie.NEW_RELEASE);
        Rental rental2 = new Rental(movie2, 5);
        customer.addRenatals(rental2);

        String expected = "Rental Record for hara\n"
            + "\t通常\t3.5\n"
            + "\t新作\t15.0\n"
            + "Amount owed is 18.5\n"
            + "You earned 3 frequent renter points";

        assertEquals(customer.statement(), expected);
    }
}
