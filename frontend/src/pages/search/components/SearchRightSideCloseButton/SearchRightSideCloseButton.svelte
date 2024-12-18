<script lang="ts">
    import Icon from "@iconify/svelte";
    import {searchFilterDrawerStore} from "@/pages/search/stores/search-filter-drawer-store.ts";

</script>

<div class="lg:hidden border-b border-light-palette-divider dark:border-dark-palette-divider h-16  flex items-center px-6">
    <button
            onclick={() => searchFilterDrawerStore.set(false)}
            class="flex items-center gap-x-3 font-semibold text-base">
        <Icon icon="hugeicons:arrow-left-02" width={18} height={18}/>
        <span>Filters</span>
    </button>
</div>