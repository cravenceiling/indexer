<script setup lang="ts">
import { useEmails } from '@/composables/useEmails'
import { computed, ref } from 'vue'
import { CalendarDaysIcon } from '@heroicons/vue/24/outline'

const { emails, error, isLoading, selectEmail } = useEmails()

const hasEmails = computed(() => emails.value.length > 0)

const sortOrder = ref<'asc' | 'desc'>('desc')

const sortByDate = () => {
  console.log('Sort by date')
  if (sortOrder.value === 'desc') {
    emails.value.sort(
      (a, b) => new Date(b._source.email.date).getTime() - new Date(a._source.email.date).getTime()
    )
    sortOrder.value = 'asc'
  } else {
    emails.value.sort(
      (a, b) => new Date(a._source.email.date).getTime() - new Date(b._source.email.date).getTime()
    )
    sortOrder.value = 'desc'
  }
}
</script>

<template>
  <div class="flex justify-center items-center w-full h-full">
    <div class="flex flex-col overflow-y-auto w-full items-center h-full">
      <div v-if="error" class="text-red-500">
        {{ error }}
      </div>
      <div v-else-if="hasEmails" class="w-full">
        <div class="flex items-center justify-between w-full">
          <!-- Buttons to sort emails -->
          <div class="flex items-center space-x-2">
            <button
              class="bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded flex items-center space-x-2"
              @click="sortByDate"
            >
              <CalendarDaysIcon class="h-5 w-5" />
              <span>Sort by Date</span>
            </button>
          </div>
        </div>
        <table class="border-separate border-spacing-y-2 w-full">
          <tbody v-for="hit in emails" :key="hit._id" class="text-gray-500">
            <tr @click="selectEmail(hit)" class="hover:bg-blue-100 hover:cursor-pointer">
              <td class="border border-gray-400 text-sm rounded-xl p-2">
                <div class="flex flex-col items-start space-y-2">
                  <span class="font-bold"> {{ hit._source.email.from }} </span>
                  <span> {{ hit._source.email.subject }} </span>
                  <span> {{ hit._source.email.date }} </span>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-else class="flex items-center justify-center h-full">
        <p class="text-gray-500">No emails found.</p>
      </div>
      <button
        v-if="hasEmails"
        class="bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded"
        :disabled="isLoading"
      >
        {{ isLoading ? 'Loading...' : 'Load More' }}
      </button>
    </div>
  </div>
</template>
