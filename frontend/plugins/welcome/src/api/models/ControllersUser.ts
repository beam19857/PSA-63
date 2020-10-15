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
     * @type {number}
     * @memberof ControllersUser
     */
    department?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersUser
     */
    doctorEmail?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersUser
     */
    doctorName?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersUser
     */
    expertise?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersUser
     */
    position?: number;
}

export function ControllersUserFromJSON(json: any): ControllersUser {
    return ControllersUserFromJSONTyped(json, false);
}

export function ControllersUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersUser {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'department': !exists(json, 'department') ? undefined : json['department'],
        'doctorEmail': !exists(json, 'doctorEmail') ? undefined : json['doctorEmail'],
        'doctorName': !exists(json, 'doctorName') ? undefined : json['doctorName'],
        'expertise': !exists(json, 'expertise') ? undefined : json['expertise'],
        'position': !exists(json, 'position') ? undefined : json['position'],
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
        
        'department': value.department,
        'doctorEmail': value.doctorEmail,
        'doctorName': value.doctorName,
        'expertise': value.expertise,
        'position': value.position,
    };
}


