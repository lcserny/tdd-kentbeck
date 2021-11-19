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

    @Test
    fun `multiplying francs is possible`() {
        val fiveFrancs = Franc(5)

        val tenFrancs = fiveFrancs.times(2)
        assertEquals(Franc(10), tenFrancs)

        val fifteenFrancs = fiveFrancs.times(3)
        assertEquals(Franc(15), fifteenFrancs)
    }

    @Test
    fun `francs can be compared`() {
        val fiveFrancs = Franc(5)
        val newFiveFrancs = Franc(5)
        assertEquals(fiveFrancs.amount, newFiveFrancs.amount)
        assertEquals(fiveFrancs, newFiveFrancs)

        val sixFrancs = Franc(6)
        val sevenFrancs = Franc(7)
        assertNotEquals(sixFrancs, sevenFrancs)
    }
}