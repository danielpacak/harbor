export class Scanner {

    scanner_id?: number;
    name?: string;
    provider?: string;
    endpointUrl?: string;
    isDefault?: boolean;
    authorization?: string;
    description?: string;

    constructor(scanner_id: number, name: string, provider: string, endpointUrl: string, isDefault: boolean) {
        this.scanner_id = scanner_id;
        this.name = name;
        this.provider = provider;
        this.endpointUrl = endpointUrl;
        this.isDefault = isDefault;
        this.authorization = 'None';
    }

}
