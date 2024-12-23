export const alphaNumericSorter = new Intl.Collator('en', {numeric: true, sensitivity: 'accent'}).compare;
export const dateSorter = (a, b) => {
    if (a.date > b.date) return 1;
    if (a.date < b.date) return -1;
    return 0;
}
export const sorterByTitle = (a, b) => {
    if (a.title > b.title) return 1;
    if (a.title < b.title) return -1;
    return 0;
}
export const sorterById = (a, b) => {
    if (a?.id > b?.id) return 1;
    if (a?.id < b?.id) return -1;
    return 0;
}
export const sorterByWebIsNotNull = (a, b) => {
    if (!!a.web > !!b.web) return 1;
    if (!!a.web < !!b.web) return -1;
    return 0;
}
export const sorterByUserLogin = (a, b) => {
    if (a.user.login === b.user.login) return 0;
    if ([a.user.login, b.user.login].sort(alphaNumericSorter) === [a.user.login, b.user.login]) {
        return -1;
    } else {
        return 1;
    }
}
export const sorterByAmount = (a, b) => {
    if (a.amount > b.amount) return -1;
    if (a.amount < b.amount) return 1;
    return 0;
}
export const getSorter = (func) => {
    return (a, b) => {
        const funcA = func(a);
        const funcB = func(b);
        if (funcA > funcB) return -1;
        if (funcA < funcB) return 1;
        return 0;
    }
}
export const getReversedSorter = (func) => {
    return (a, b) => {
        const funcA = func(a);
        const funcB = func(b);
        if (funcA > funcB) return 1;
        if (funcA < funcB) return -1;
        return 0;
    }
}

export const reverseSorter = (sorter) => {
    return (a, b) => -sorter(a, b)
}
