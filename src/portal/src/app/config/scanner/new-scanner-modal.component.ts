import {Component, Output, EventEmitter, ViewChild} from '@angular/core';
import {Scanner, ScannerMetadata} from "./scanner";
import {NewScannerFormComponent} from "./new-scanner-form.component";
import {ConfigScannerService} from "./config-scanner.service";

@Component({
    selector: "new-scanner-modal",
    templateUrl: "new-scanner-modal.component.html",
    styleUrls: ['../../common.scss']
})
export class NewScannerModalComponent {

    opened: boolean = false;
    scannerMetadata: ScannerMetadata;

    @Output() notify = new EventEmitter<Scanner>();

    @ViewChild(NewScannerFormComponent)
    newScannerForm: NewScannerFormComponent;

    constructor(
        private configScannerService: ConfigScannerService
    ) {
    }

    open(scanner: Scanner): void {
        delete this.scannerMetadata;
        this.newScannerForm.setData(scanner);
        this.opened = true;
        if (scanner.id) {
            this.configScannerService.getMetadata(scanner.id).subscribe(metadata => {
                this.scannerMetadata = metadata;
            }, error => {
                delete this.scannerMetadata;
                alert('Error while getting metadata: ' + JSON.stringify(error));
            });
        }
    }

    close(): void {
        this.opened = false;
    }

    testConnection(): void {
        this.configScannerService.ping(this.newScannerForm.getData().endpoint_url).subscribe(metadata => {
            this.scannerMetadata = metadata;
        }, error => {
            delete this.scannerMetadata;

            alert("Connection failed: " + JSON.stringify(error));
        });
    }

    ok(): void {
        let scanner = this.newScannerForm.getData();
        if (scanner.id) {
            this.configScannerService.update(scanner).subscribe(() => {
                this.notify.emit(scanner);
                this.close();
            }, error => {
                alert("Error while updating scanner registration: " + error);
            });
        } else {
            this.configScannerService.create(scanner).subscribe(() => {
                this.notify.emit(scanner);
                this.close();
            }, error => {
                alert("Error while creating scanner registration: " + error);
            });
        }
    }

}
