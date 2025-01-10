<script lang="ts">
    type TextAreaProps = {
        placeholder: string,
        styles?: string,
    };

    const {placeholder, styles}: TextAreaProps = $props();
</script>

<textarea
        placeholder={placeholder}
        class="{styles} placeholder:text-light-palette-text-secondary placeholder:dark:text-dark-palette-text-secondary w-full h-40 p-3 bg-light-palette-background-block dark:bg-dark-palette-background-block text-base ring-1 ring-light-palette-divider dark:ring-dark-palette-divider rounded-md focus:ring-cyan-500 focus:dark:ring-cyan-500 outline-none resize-none"
></textarea>