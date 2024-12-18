<script lang="ts">
    import SearchFromToBlockInputSpinnersLayout
        from "@/pages/search/components/SearchFromToBlockInputSpinners/SearchFromToBlockInputSpinnersLayout.svelte";
    import SearchFromToBlockInputSpinnersIcon
        from "@/pages/search/components/SearchFromToBlockInputSpinnersIcon/SearchFromToBlockInputSpinnersIcon.svelte";

    type SearchFromToBlockInputProps = {
        value: number,
        setValue: (value: number) => void,
    }

    const {value, setValue}: SearchFromToBlockInputProps = $props();
</script>

<SearchFromToBlockInputSpinnersLayout>
    <SearchFromToBlockInputSpinnersIcon
            value={value}
            setValue={setValue}
            direction="up"
    />
    <SearchFromToBlockInputSpinnersIcon
            value={value}
            setValue={setValue}
            direction="down"
    />
</SearchFromToBlockInputSpinnersLayout>