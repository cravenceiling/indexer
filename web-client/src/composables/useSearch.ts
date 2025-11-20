import { EmailResponse } from '@/models/EmailResponse'

export async function searchEmails(term: string) {
  const API_URL = import.meta.env.API_URL ?? 'http://localhost:8080'

  const url = new URL(`${API_URL}/emails/search`)
  url.searchParams.append('index', 'enron')
  url.searchParams.append('page', '10')
  url.searchParams.append('term', term)

  const response = await fetch(url)

  if (!response.ok) {
    console.error('Error in request: ', response)
    return {
      error: 'There was an error searching for emails, please try again.',
      data: []
    }
  }

  const result = (await response.json()) as EmailResponse
  const data = result.hits.hits
  console.log('data: ', data)

  return { data, error: null }
}
