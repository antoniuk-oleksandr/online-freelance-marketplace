export const getPaginationRange = (
  currentPage: number, totalPages: number
): (string | number)[] => {
  if (totalPages <= 5) {
    return Array.from({ length: totalPages }, (_, i) => i + 1)
  }

  if (totalPages > 4 && currentPage <= 4) {
    return [...Array.from({ length: 5 }, (_, i) => i + 1), "...", totalPages]
  }

  if (totalPages > 4 && currentPage > 4 && currentPage < totalPages - 3) {
    return [1, "...", ...Array.from({ length: 3 }, (_, i) => i + currentPage - 1), "...", totalPages]
  }


  return [1, "...", ...Array.from({ length: 5 }, (_, i) => totalPages + i - 4)]
}
