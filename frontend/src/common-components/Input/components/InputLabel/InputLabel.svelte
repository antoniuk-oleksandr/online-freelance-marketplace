<script lang="ts">
    type InputLabelProps = {
        id: string,
        value: string,
        label: string,
        error: string | null,
        wasSubmitted: boolean,
        isFocused: boolean,
    }

    const {id, value, label, wasSubmitted, isFocused, error}: InputLabelProps = $props();
</script>

<label
        for={id}
        class="{isFocused || value.length > 0 ? '-translate-y-[32.5px] text-sm top-[21px]' : 'text-lg top-[16px]'}
                {wasSubmitted && error ? 'text-red-500' : isFocused ? 'text-cyan-500' : 'text-light-palette-divider dark:text-dark-palette-divider'}
                duration-translate cursor-text bg-light-palette-background-block dark:bg-dark-palette-background-block px-1 font-medium absolute left-3"
>{label}</label>