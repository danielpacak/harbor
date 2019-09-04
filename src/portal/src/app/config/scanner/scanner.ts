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

export class ScannerMetadata {
    name: string;
    vendor: string;
    version: string;
}
