<script lang="ts">
    import Icon from "@iconify/svelte";
    import {searchFilterArrayStore} from "@/pages/search/stores/search-filter-array-store.ts";
</script>

<button
        onclick={() => searchFilterArrayStore.update((prev) => ({...prev, isOpened: false}))}
        class="flex gap-x-3 items-center">
    <Icon icon="hugeicons:arrow-left-02" width={18} height={18}/>
    <span class="">Back</span>
</button>