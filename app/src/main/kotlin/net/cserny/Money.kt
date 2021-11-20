package net.cserny

abstract class Money(open val amount: Int) {
    abstract fun times(multiplier: Int): Money
}

data class Dollar(override val amount: Int) : Money(amount) {

    override fun times(multiplier: Int): Money {
        return Dollar(amount * multiplier)
    }
}

data class Franc(override val amount: Int) : Money(amount) {

    override fun times(multiplier: Int): Money {
        return Franc(amount * multiplier)
    }
}