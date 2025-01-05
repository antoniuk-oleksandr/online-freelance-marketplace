<script lang="ts">
    import SearchFromToBlockInput from "@/pages/search/components/SearchFromToBlockInput/SearchFromToBlockInput.svelte";
    import Icon from "@iconify/svelte";
    import SearchFromToBlockLayout from "@/pages/search/components/SearchFromToBlock/SearchFromToBlockLayout.svelte";

    type SearchFromToBlockProps = {
        title: string,
        from: number,
        setFrom: (value: number) => void,
        to: number,
        setTo: (value: number) => void,
    }

    const {title, from, setFrom, to, setTo}: SearchFromToBlockProps = $props();
</script>

<SearchFromToBlockLayout>
    <div class="font-semibold">{title}</div>
    <div class="flex items-center gap-x-3">
        <SearchFromToBlockInput
                type="From"
                value={from}
                setValue={setFrom}
        />
        <Icon icon="hugeicons:minus-sign" width={16} height={16}/>
        <SearchFromToBlockInput
                type="To"
                value={to}
                setValue={setTo}
        />
    </div>
</SearchFromToBlockLayout>