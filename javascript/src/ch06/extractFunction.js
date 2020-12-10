function printOwing(invoice) {
    let outstatnding = 0;

    printBanner();

    for (const o of invoice.orders) {
        outstatnding += o.amount;
    }

    const today = Clock.today;
    invoice.dueDate = new Date(today.getFullYear(), today.getMonth(), today.getDate() + 30);

    console.log(`name: ${invoice.customer}`);
    console.log(`amount: ${outstanding}`);
    console.log(`due: ${invoice.dueDate.toLocalDateString()}`);

    function printBanner() {
        console.log("******************************")
        console.log("******* Customer Owes *******")
        console.log("******************************")
    }
}