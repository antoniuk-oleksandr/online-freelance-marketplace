<script lang="ts">
    import {searchFilterDrawerStore} from "@/pages/search/stores/search-filter-drawer-store";
</script>

<button
        aria-label="Close filters modal"
        onclick={() => searchFilterDrawerStore.set(false)}
        class="bg-black !duration-100 bg-opacity-50 w-full h-full lg:hidden"
></button>
