#[cfg(test)]
mod tests {
    use super::money::*;

    #[test]
    fn multiplying_dollars_is_possible() {
        let five_dollars = Dollar::new(5);

        let ten_dollars = five_dollars.times(2);
        assert_eq!(Dollar::new(10), ten_dollars);

        let fifteen_dollars = five_dollars.times(3);
        assert_eq!(Dollar::new(15), fifteen_dollars);
    }

    #[test]
    fn dollars_can_be_compared() {
        let five = Dollar::new(5);
        let new_five = Dollar::new(5);
        assert_eq!(five.amount, new_five.amount);
        assert_eq!(five, new_five);

        let six = Dollar::new(6);
        let seven = Dollar::new(7);
        assert_ne!(six, seven);
    }

    #[test]
    fn multiplying_francs_is_possible() {
        let five_francs = Franc::new(5);

        let ten_francs = five_francs.times(2);
        assert_eq!(Franc::new(10), ten_francs);

        let fifteen_francs = five_francs.times(3);
        assert_eq!(Franc::new(15), fifteen_francs);
    }

    #[test]
    fn francs_can_be_compared() {
        let five = Franc::new(5);
        let new_five = Franc::new(5);
        assert_eq!(five.amount, new_five.amount);
        assert_eq!(five, new_five);

        let six = Franc::new(6);
        let seven = Franc::new(7);
        assert_ne!(six, seven);
    }
}

pub mod money {
    pub trait MoneyOperations {
        fn times(&self, multiplier: u32) -> Self;
    }

    #[derive(PartialEq, Debug)]
    pub struct Dollar {
        pub amount: u32,
    }

    impl Dollar {
        pub fn new(amount: u32) -> Dollar {
            Dollar { amount }
        }
    }

    impl MoneyOperations for Dollar {
        fn times(&self, multiplier: u32) -> Dollar {
            Dollar { amount: self.amount * multiplier }
        }
    }

    #[derive(PartialEq, Debug)]
    pub struct Franc {
        pub amount: u32,
    }

    impl Franc {
        pub fn new(amount: u32) -> Franc {
            Franc { amount }
        }
    }

    impl MoneyOperations for Franc {
        fn times(&self, multiplier: u32) -> Franc {
            Franc { amount: self.amount * multiplier }
        }
    }
}

