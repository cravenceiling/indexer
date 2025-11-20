import { EmailResponse } from '@/models/EmailResponse';
import type { Hit } from '@/models/Hit';
import { ref } from 'vue';

const emails = ref<Hit[]>([]);
const isLoading = ref(false);
const error = ref<string | null>(null);
const selectedEmail = ref<Hit | null>(null);
const term = ref('');

export function useEmails() {
  const searchEmails = async (page: number, limit: number) => {
    const API_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:8080';

    isLoading.value = true;
    error.value = null;

    try {
      const url = new URL(`${API_URL}/emails/search`);
      url.searchParams.append('index', 'enron');
      url.searchParams.append('term', encodeURIComponent(term.value));
      url.searchParams.append('page', (page * limit).toString());
      url.searchParams.append('limit', limit.toString());
      const response = await fetch(url);

      if (!response.ok) {
        throw new Error(`API error: ${response.status}`);
      }

      const result = (await response.json()) as EmailResponse;
      const data: Hit[] = result.hits.hits;
      emails.value = data;
    } catch (err: unknown) {
      console.error('Error searching for emails: ', err);
      error.value = 'There was an error searching for emails, please try again.';
    } finally {
      isLoading.value = false;
    }
  };

  const selectEmail = (email: Hit) => {
    selectedEmail.value = email;
  };

  return {
    emails,
    searchEmails,
    isLoading,
    error,
    selectedEmail,
    selectEmail,
    term,
  };
}
