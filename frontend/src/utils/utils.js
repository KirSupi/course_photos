export const formatBookingHours = h => {
    return  (h < 10 ? '0' : '') + h + ':00-' + ((h + 1) % 24 < 10 ? '0' : '') + ((h + 1) % 24) + ':00';
}