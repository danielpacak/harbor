import {Component, Output, EventEmitter, ViewChild} from '@angular/core';
import {Scanner} from "./scanner";
import {NewScannerFormComponent} from "./new-scanner-form.component";
import {ConfigScannerService} from "./config-scanner.service";

@Component({
    selector: "new-scanner-modal",
    templateUrl: "new-scanner-modal.component.html",
    styleUrls: ['../../common.scss']
})
export class NewScannerModalComponent {

    opened: boolean = false;

    @Output() notify = new EventEmitter<Scanner>();

    @ViewChild(NewScannerFormComponent)
    newScannerForm: NewScannerFormComponent;

    constructor(
        private configScannerService: ConfigScannerService
    ) {}

    open(): void {
        this.opened = true;
    }

    close(): void {
        this.opened = false;
    }

    create(): void {
        // TODO Save it
        let s = this.newScannerForm.getData();
        this.configScannerService.create(s);
        this.notify.emit(s);
        // show mesae
    }

}
