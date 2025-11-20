import { EmailResponse } from '@/models/EmailResponse'
import type { Hit } from '@/models/Hit'
import { ref } from 'vue'

const emails = ref<Hit[]>([])
const isLoading = ref(false)
const error = ref<string | null>(null)
const selectedEmail = ref<Hit | null>(null)

export function useEmails() {
  const searchEmails = async (search: string) => {
    const API_URL = import.meta.env.API_URL ?? 'http://localhost:8080'

    const term = search.trim()
    isLoading.value = true
    error.value = null

    const url = new URL(`${API_URL}/emails/search`)
    url.searchParams.append('index', 'enron')
    url.searchParams.append('page', '10')
    url.searchParams.append('term', encodeURIComponent(term))

    try {
      const response = await fetch(url)

      if (!response.ok) {
        throw new Error(`API error: ${response.status}`)
      }

      const result = (await response.json()) as EmailResponse
      const data: Hit[] = result.hits.hits
      emails.value = data
    } catch (err: unknown) {
      console.error('Error searching for emails: ', err)
      error.value = 'There was an error searching for emails, please try again.'
    } finally {
      isLoading.value = false
    }
  }

  const selectEmail = (email: Hit) => {
    selectedEmail.value = email
  }

  return {
    emails,
    searchEmails,
    isLoading,
    error,
    selectedEmail,
    selectEmail
  }
}
