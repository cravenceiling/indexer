<script setup lang="ts">
import { useEmails } from '@/composables/useEmails';

const { searchEmails, term } = useEmails();

const handleSearch = async () => {
  term.value = term.value.trim();
  await searchEmails(0, 10);
};
</script>

<template>
  <div class="items-center w-full min-h-10 rounded-lg bg-white overflow-hidden">
    <form @submit.prevent="handleSearch" class="flex items-center w-full h-auto min-h-10 px-0 py-0">
      <input
        type="search"
        placeholder="Search..."
        value="term"
        v-model="term"
        class="flex-1 bg-transparent px-3 sm:px-4 py-2 text-gray-700 placeholder-gray-400 focus:outline-none min-w-0"
      />

      <button
        type="submit"
        class="flex items-center justify-center px-3 sm:px-4 min-h-10 text-white bg-blue-600 hover:bg-blue-700 rounded-r-lg transition-colors"
      >
        <svg
          class="w-4 sm:w-5"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          />
        </svg>
      </button>
    </form>
  </div>
</template>

<style></style>
