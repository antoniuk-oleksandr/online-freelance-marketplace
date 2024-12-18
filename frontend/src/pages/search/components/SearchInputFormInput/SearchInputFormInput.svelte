<script lang="ts">
    import {handleSearchInput} from "@/pages/search/handlers.ts";

    type SearchInputFormInputProps =  {
        query: string | null,
    }

    const {query}: SearchInputFormInputProps = $props();
</script>

<input
        placeholder="Search for services"
        oninput={handleSearchInput}
        value={query}
        class="outline-none placeholder:text-light-palette-text-secondary placeholder:dark:text-dark-palette-text-secondary w-full bg-transparent rounded-md text-base px-3 py-2 ring-1 ring-light-palette-divider dark:ring-dark-palette-divider focus:ring-cyan-500 focus:dark:ring-cyan-500 duration-200 ease-out"
/>