<script setup lang="ts">
import AppHeader from './components/AppHeader.vue';
import ContentPane from './components/ContentPane.vue';
import InboxTable from './components/InboxTable.vue';
</script>

<template>
  <div class="h-screen flex flex-col">
    <AppHeader />

    <main class="flex-1 w-full p-4 sm:p-6 lg:p-8 space-y-6 overflow-y-auto">
      <div class="flex flex-col md:flex-row w-full h-full">
        <div class="w-full md:w-1/2 h-full border border-gray-400 rounded-md p-4">
          <InboxTable />
        </div>

        <div class="w-full md:w-1/2 h-full p-4 border border-gray-400 rounded-md">
          <ContentPane />
        </div>
      </div>
    </main>
  </div>
</template>

<style></style>
