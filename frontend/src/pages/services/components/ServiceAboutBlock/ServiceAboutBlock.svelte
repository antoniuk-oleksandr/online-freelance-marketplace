<script lang="ts">
    import ServiceAboutBlockLayout from "@/pages/services/components/ServiceAboutBlock/ServiceAboutBlockLayout.svelte";
    import type {Service} from "@/types/Service.ts";

    const {description}: Service = $props();
</script>

<ServiceAboutBlockLayout>
    <h3 class="text-xl font-bold text-light-palette-text-primary dark:text-dark-palette-text-primary">About</h3>
    <p
            class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
    >{description}</p>
</ServiceAboutBlockLayout>