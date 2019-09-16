// TODO Rename to ScannerRegistration -> We call it like that in the proposal
export class Scanner {

    id?: number;
    name?: string;
    provider?: string;
    endpoint_url?: string;
    default?: boolean;
    enabled: boolean;
    deleted: boolean;
    authorization?: string;
    description?: string;
}

export class ScannerProperties {
    name: string;
    vendor: string;
    version: string;
}

export class ScannerMetadata {
    scanner: ScannerProperties;
}
