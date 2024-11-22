<script lang="ts">
    import Icon from "@iconify/svelte";
    import {handleToastCloseButtonClick} from "@/common-components/ToastElement/handlers.ts";
</script>

<button
        class="ml-4 text-2xl text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
        onclick={handleToastCloseButtonClick}>
    <Icon icon="material-symbols:close-small" />
</button>