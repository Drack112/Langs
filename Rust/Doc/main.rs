fn main(){
    /// Line comments; document the next item
    /** Block comments; document the next item */

    //! Line comments; document the enclosing item
    /*! Block comments; document the enclosing item !*/
    
    /// Outer comment
    #[doc = "Outer comment"]

    //! Inner comment
    #![doc = "Inner comment"]
    println!("Hello World!")
}