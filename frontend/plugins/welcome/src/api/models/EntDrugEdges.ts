/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntRequisition,
    EntRequisitionFromJSON,
    EntRequisitionFromJSONTyped,
    EntRequisitionToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDrugEdges
 */
export interface EntDrugEdges {
    /**
     * Requisitions holds the value of the requisitions edge.
     * @type {Array<EntRequisition>}
     * @memberof EntDrugEdges
     */
    requisitions?: Array<EntRequisition>;
}

export function EntDrugEdgesFromJSON(json: any): EntDrugEdges {
    return EntDrugEdgesFromJSONTyped(json, false);
}

export function EntDrugEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDrugEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'requisitions': !exists(json, 'requisitions') ? undefined : ((json['requisitions'] as Array<any>).map(EntRequisitionFromJSON)),
    };
}

export function EntDrugEdgesToJSON(value?: EntDrugEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'requisitions': value.requisitions === undefined ? undefined : ((value.requisitions as Array<any>).map(EntRequisitionToJSON)),
    };
}


