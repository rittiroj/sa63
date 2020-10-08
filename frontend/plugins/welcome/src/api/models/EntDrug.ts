/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
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
    EntDrugEdges,
    EntDrugEdgesFromJSON,
    EntDrugEdgesFromJSONTyped,
    EntDrugEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDrug
 */
export interface EntDrug {
    /**
     * DrugsName holds the value of the "drugsName" field.
     * @type {string}
     * @memberof EntDrug
     */
    drugsName?: string;
    /**
     * 
     * @type {EntDrugEdges}
     * @memberof EntDrug
     */
    edges?: EntDrugEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDrug
     */
    id?: number;
    /**
     * Value holds the value of the "value" field.
     * @type {number}
     * @memberof EntDrug
     */
    value?: number;
}

export function EntDrugFromJSON(json: any): EntDrug {
    return EntDrugFromJSONTyped(json, false);
}

export function EntDrugFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDrug {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'drugsName': !exists(json, 'drugsName') ? undefined : json['drugsName'],
        'edges': !exists(json, 'edges') ? undefined : EntDrugEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'value': !exists(json, 'value') ? undefined : json['value'],
    };
}

export function EntDrugToJSON(value?: EntDrug | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'drugsName': value.drugsName,
        'edges': EntDrugEdgesToJSON(value.edges),
        'id': value.id,
        'value': value.value,
    };
}


