function printOwing(invoice) {
    printBanner();
    const outstatnding = calculateOutstanding(invoice);

    function printBanner() {
        console.log("******************************")
        console.log("******* Customer Owes *******")
        console.log("******************************")
    }
    recordDueDate(invoice);
    printDetails(invoice, outstatnding);
}

function calculateOutstanding(invoice) {
    let result = 0;
    for (const o of invoice.orders) {
        result += o.amount;
    }
    return result;
}

function recordDueDate(invoice) {
    const today = Clock.today;
    invoice.dueDate = new Date(today.getFullYear(), today.getMonth(), today.getDate() + 30);
}

function printDetails(invoice, outstatnding) {
    console.log(`name: ${invoice.customer}`);
    console.log(`amount: ${outstatnding}`);
    console.log(`due: ${invoice.dueDate.toLocalDateString()}`);
}