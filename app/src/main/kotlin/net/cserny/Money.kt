package net.cserny

data class Dollar(val amount: Int) {

    fun times(multiplier: Int): Dollar {
        return Dollar(amount * multiplier)
    }
}

data class Franc(val amount: Int) {

    fun times(multiplier: Int): Franc {
        return Franc(amount * multiplier)
    }
}