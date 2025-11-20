<script setup lang="ts">
import { useEmails } from '@/composables/useEmails';

const { selectedEmail } = useEmails();
</script>

<template>
  <div class="overflow-y-auto h-full">
    <div class="text-sm h-full">
      <div v-if="selectedEmail" class="flex flex-col h-full w-full space-y-4">
        <p class=""><strong>Subject:</strong> {{ selectedEmail._source.email.subject }}</p>
        <p class=""><strong>From:</strong> {{ selectedEmail._source.email.from }}</p>
        <p class=""><strong>To:</strong> {{ selectedEmail._source.email.to }}</p>
        <p class=""><strong>Date:</strong> {{ selectedEmail._source.email.date }}</p>
        <div class="px-20"></div>
        <p class="text-sm whitespace-pre-wrap">
          {{ selectedEmail._source.email.body }}
        </p>
      </div>
      <div v-else class="flex items-center justify-center text-center h-full">
        <p class="text-gray-500 dark:text-gray-400 text-lg">
          The email content will be displayed here.
        </p>
      </div>
    </div>
  </div>
</template>
