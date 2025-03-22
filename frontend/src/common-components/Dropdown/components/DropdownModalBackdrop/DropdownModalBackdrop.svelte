<script lang="ts">
    import { fade } from "svelte/transition"

</script>

<!-- svelte-ignore element_invalid_self_closing_tag -->
<div
  transition:fade={{ duration: 300 }}
  class="p-6 !z-40 flex-col fixed top-0 left-0 size-full lg:hidden bg-black bg-opacity-50 flex justify-center items-center"
/>
