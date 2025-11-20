<script setup lang="ts">
import SearchBar from './SearchBar.vue';
</script>

<template>
  <header class="bg-[#11009E] text-white">
    <div class="flex flex-col sm:flex-row items-center justify-center gap-4 sm:gap-8 p-4">
      <div class="flex items-center gap-3 sm:gap-4">
        <img
          src="../assets/mail-box.png"
          alt="ZincMail logo"
          class="w-8 h-8 sm:w-10 sm:h-10 object-contain"
        />
        <h1 class="text-2xl font-semibold">ZincMail</h1>
      </div>

      <div class="w-full sm:flex-1 max-w-2xl">
        <SearchBar />
      </div>
    </div>
  </header>
</template>
