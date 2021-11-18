package net.cserny

import org.junit.Test
import kotlin.test.assertEquals
import kotlin.test.assertNotEquals

class MoneyTest {

    @Test
    fun `multiplying dollars is possible`() {
        val fiveDollars = Dollar(5)

        val tenDollars = fiveDollars.times(2)
        assertEquals(Dollar(10), tenDollars)

        val fifteenDollars = fiveDollars.times(3)
        assertEquals(Dollar(15), fifteenDollars)
    }

    @Test
    fun `dollars can be compared`() {
        val fiveDollars = Dollar(5)
        val newFiveDollars = Dollar(5)
        assertEquals(fiveDollars.amount, newFiveDollars.amount)
        assertEquals(fiveDollars, newFiveDollars)

        val sixDollars = Dollar(6)
        val sevenDollars = Dollar(7)
        assertNotEquals(sixDollars, sevenDollars)
    }
}