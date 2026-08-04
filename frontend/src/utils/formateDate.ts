const dateFormatter = new Intl.DateTimeFormat('ru-RU', {
  day: 'numeric',
  month: 'long',
  year: 'numeric',
})

export function formatDate(value: string) {
  return dateFormatter.format(new Date(`${value}T00:00:00`))
}

export function formatDateSince(value: string) {
  return `На Авито с ${new Date(`${value}T00:00:00`).getFullYear()} года`
}
