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
/**
 * 
 * @export
 * @interface ControllersUser
 */
export interface ControllersUser {
    /**
     * 
     * @type {string}
     * @memberof ControllersUser
     */
    doctorEmail?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersUser
     */
    doctorID?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersUser
     */
    doctorName?: string;
}

export function ControllersUserFromJSON(json: any): ControllersUser {
    return ControllersUserFromJSONTyped(json, false);
}

export function ControllersUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersUser {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'doctorEmail': !exists(json, 'doctorEmail') ? undefined : json['doctorEmail'],
        'doctorID': !exists(json, 'doctorID') ? undefined : json['doctorID'],
        'doctorName': !exists(json, 'doctorName') ? undefined : json['doctorName'],
    };
}

export function ControllersUserToJSON(value?: ControllersUser | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'doctorEmail': value.doctorEmail,
        'doctorID': value.doctorID,
        'doctorName': value.doctorName,
    };
}

