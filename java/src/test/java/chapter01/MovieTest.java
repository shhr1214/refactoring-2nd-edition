package chapter01;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;

class MovieTests {
    @Test
    @DisplayName("1 + 1 = 2")
    void addTwoNumbers() {
        assertEquals(2, 1 + 1, "1 + 1 should equal 2");
    }

    @ParameterizedTest(name = "{0} + {1} = {2}")
}
