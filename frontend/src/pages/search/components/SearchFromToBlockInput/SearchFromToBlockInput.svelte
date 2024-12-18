<script lang="ts">
    import SearchFromToBlockInputSpinners
        from "@/pages/search/components/SearchFromToBlockInputSpinners/SearchFromToBlockInputSpinners.svelte";

    type SearchFromToBlockInputProps = {
        value: number,
        setValue: (value: number) => void,
        type: "From" | "To",
    }

    const {value, setValue, type}: SearchFromToBlockInputProps = $props();
</script>

<div class="w-1/2 relative">
    <input
            class="w-full placeholder:text-light-palette-text-secondary dark:placeholder:text-dark-palette-text-secondary duration-200 ease-out bg-transparent ring-1 ring-light-palette-divider dark:ring-dark-palette-divider dark:focus:ring-cyan-500 focus:ring-cyan-500 outline-none rounded-md py-1 px-3"
            value={Number.isNaN(value) ? "" : value}
            onchange={(event) => setValue(parseInt(event.currentTarget.value))}
            type="number"
            placeholder={type}
    />
    <SearchFromToBlockInputSpinners
        value={value}
        setValue={setValue}
    />
</div>
