<script lang="ts">
    import Input from "@/common-components/Sign/components/SignInput/Input.svelte";
    import type {FormStore} from "@/types/FormStore.ts";
    import {formStore} from "@/common-components/Sign/stores/form-store.ts";

    let formStoreData = $state<null | FormStore>(null);
    formStore.subscribe((value) => formStoreData = value);
</script>

{#if formStoreData && formStoreData.data && formStoreData.errors}
    <Input
            wasSubmitted={formStoreData.wasSubmitted}
            value={formStoreData.data.usernameOrEmail}
            error={formStoreData.errors.usernameOrEmail?.[0] || null}
            id="usernameOrEmail"
            type="text"
            label="Username or Email"
    />
{/if}
